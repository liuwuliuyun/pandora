package dataapigenerator

import (
	"fmt"
	"log"
	"path"
	"strings"
)

func (s Service) generateVersionDefinition() error {
	if s.debugLog {
		log.Printf("[DEBUG] Checking for an existing Generation Settings file..")
	}

	// recreate the directory
	excludeList := []string{
		"ApiVersionDefinition-GenerationSetting.cs",
	}
	if err := recreateDirectoryExcludingFiles(s.WorkingDirectoryForApiVersion, excludeList, s.debugLog); err != nil {
		return fmt.Errorf("recreating directory %q: %+v", s.WorkingDirectoryForApiVersion, err)
	}

	// then generate the files
	if s.debugLog {
		log.Printf("[DEBUG] Generating Api Version Definition..")
	}
	isPreview := strings.Contains(strings.ToLower(s.data.ApiVersion), "preview")
	definitionFilePath := path.Join(s.WorkingDirectoryForApiVersion, "ApiVersionDefinition.cs")
	if err := writeToFile(definitionFilePath, codeForApiVersionDefinition(s.NamespaceForApiVersion, s.data.ApiVersion, isPreview, s.data.Resources)); err != nil {
		return fmt.Errorf("writing Api Version Definition to %q: %+v", definitionFilePath, err)
	}

	generationSettingFilePath := path.Join(s.WorkingDirectoryForApiVersion, "ApiVersionDefinition-GenerationSetting.cs")
	exists, err := fileExistsAtPath(generationSettingFilePath)
	if err != nil {
		return err
	}
	if *exists {
		if s.debugLog {
			log.Printf("[DEBUG] Api Version Definition Generation Settings exists at %q - skipping", generationSettingFilePath)
		}
		return nil
	}

	if s.debugLog {
		log.Printf("[DEBUG] Generating Api Version Definition Generation Setting..")
	}
	if err := writeToFile(generationSettingFilePath, codeForApiVersionDefinitionSetting(s.NamespaceForApiVersion)); err != nil {
		return fmt.Errorf("writing Api Version Definition Generation Setting to %q: %+v", definitionFilePath, err)
	}

	return nil
}
