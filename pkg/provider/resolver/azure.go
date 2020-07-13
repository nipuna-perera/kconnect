package resolver

import (
	"fmt"

	"github.com/fidelity/kconnect/pkg/provider/identity"
	//"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type AzureFlagsResolver struct {
}

func (r *AzureFlagsResolver) Resolve(identity identity.Identity, _ *pflag.FlagSet) error {
	fmt.Println("In AzureFlagsResolver.Resolve\n")

	//TODO get client from identity

	//resourceGroup := viper.GetString("resource-group")
	resourceGroup := ""

	if resourceGroup == "" {
		if err := r.resolveResourceGroup(); err != nil {
			return fmt.Errorf("failed to resolve resource group: %w", err)
		}
	}

	return nil
}

func (r *AzureFlagsResolver) resolveResourceGroup() error {
	//TODO: azure client will be accessible
	//TODO: query the azure API

	// NOTE: testing creating a very long list
	/*resourceGroups := []string{}
	for i := 0; i < 100; i++ {
		resourceGroupName := fmt.Sprintf("reg-test-%d", i)
		resourceGroups = append(resourceGroups, resourceGroupName)
	}

	prompt := promptui.Select{
		Label: "Resource group",
		Items: resourceGroups,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("failed prompting for resource group: %w", err)
	}*/
	result := "rg-dev"

	settings := viper.AllSettings()
	for k, v := range settings {
		fmt.Printf("%s = %s\n", k, v)
	}

	viper.SetDefault("resource-group", result)

	return nil
}
