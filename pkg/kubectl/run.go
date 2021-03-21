package kubectl

import (
	"fmt"
	"io/ioutil"
	"log"
	"multi-kubectl/pkg/contexts"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//Opts are Options for the contexts configuration
var Opts struct {
	Contexts []*string `long:"contexts" description:"k8s contexts to use"`
}

//RunCommand runs the kubectl command in parallel for all configured contexts
func RunCommand(args []string) {
	contextsFromFile, err := contexts.NewKubeConfigFromFile(getKubeconfigPath())
	args, ctxs := fetchContexts(contextsFromFile, args)

	if err != nil {
		log.Fatal("error loading contexts from kubeconfig", err)
	}

	var wg sync.WaitGroup
	for _, ctx := range ctxs {

		if !contextsFromFile.ContextExists(ctx) {
			fmt.Printf("Skipping context %s as cannot be found in KUBECONFIG file\n", ctx)
			continue
		}
		wg.Add(1)
		go runCommand(ctx, args, &wg)
	}
	wg.Wait()

}

func runCommand(ctx string, args []string, wg *sync.WaitGroup) {
	defer wg.Done()
	tmpFile, err := ioutil.TempFile(os.TempDir(), fmt.Sprintf("kconfig-%s", ctx))
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())

	execCmd("kubectl", args, getKubeconfigPath(), ctx, false)
}

func execCmd(execCmd string, args []string, configPath string, ctx string, silent bool) error {
	args = append(args, "--context", ctx)
	execPath, _ := exec.LookPath(execCmd)
	args = append([]string{execCmd}, args...)
	krun := &exec.Cmd{
		Path: execPath,
		Args: args,
	}
	krun.Env = os.Environ()
	krun.Env = append(krun.Env, fmt.Sprintf("KUBECONFIG=%s", configPath))
	output, err := krun.CombinedOutput()

	if !silent {
		fmt.Println(fmt.Sprintf("cluster:%s\n%s", ctx, string(output)))
	}
	if err != nil {
		fmt.Println(string(output))
		fmt.Printf("%+v\n", err)
	}
	return err
}

func getKubeconfigPath() string {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding homedir, set KUBECONFIG env to avoid auto-detect: %+v", err)
			os.Exit(1)
		}
		kubeconfig = fmt.Sprintf("%s/.kube/config", dir)
	}
	return kubeconfig
}

func fetchContexts(kcontexts *contexts.KubeConfig, a []string) (args []string, contexts []string) {
	for i := 1; i < len(a); i++ {
		if a[i] == "--contexts" || a[i] == "-contexts" || a[i] == "--ctx" || a[i] == "-ctx" {
			if i < len(a)-1 {
				contexts = append(contexts, a[i+1])
				i++
			}
			continue
		}
		if strings.HasPrefix(a[i], "--contexts=") || strings.HasPrefix(a[i], "-contexts=") ||
			strings.HasPrefix(a[i], "--ctx=") || strings.HasPrefix(a[i], "-ctx=") {
			contexts = append(contexts, strings.Split(a[i], "=")[1])
			continue
		}

		if a[i] == "--match-contexts" || a[i] == "-match-contexts" || a[i] == "--match-ctx" || a[i] == "-match-ctx" {
			if i < len(a)-1 {
				matched := kcontexts.GetMatchedContexts(a[i+1])
				contexts = append(contexts, matched...)
				i++
			}
			continue
		}
		if strings.HasPrefix(a[i], "--match-contexts=") || strings.HasPrefix(a[i], "-match-contexts=") ||
			strings.HasPrefix(a[i], "--match-ctx=") || strings.HasPrefix(a[i], "-match-ctx=") {
			matched := kcontexts.GetMatchedContexts(strings.Split(a[i], "=")[1])
			contexts = append(contexts, matched...)
			continue
		}
		args = append(args, a[i])
	}
	return
}
