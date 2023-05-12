package files

import (
	"fmt"
	"strings"
)

/** Create Dto file*/
func CreateDtoFile(titledName, name string, auth, documentation, logger bool, properties []string) [2]string {

	imports := ""
	headerContent := fmt.Sprintf("export class Create%sDto {", titledName)
	classValidatorImports := []string{"IsOptional", "IsNotEmpty"}
	content := ""

	if documentation {
		imports += "import { ApiProperty } from '@nestjs/swagger';"
	}

	for _, property := range properties {
		if property != "" {

			colonIndex := strings.IndexRune(property, ':')
			optional := strings.ContainsRune(property, '?')
			keyName := property[:colonIndex]
			keyValue := property[colonIndex+1:]

			if documentation {
				optionalTag := ""
				if optional {
					optionalTag = ", nullable: true"
				}

				content += fmt.Sprintf("\t@ApiProperty({description: '%s %s' %s })\n", titledName, keyName, optionalTag)
			}

			if optional {
				content += "\t@IsOptional()\n"
			} else {
				content += "\t@IsNotEmpty()\n"
			}

			if keyValue == "string" {
				content += "\t@IsString()\n"
			} else if keyValue == "number" {
				content += "\t@IsInt()\n"
			} else if keyValue == "email" {
				content += "\t@IsEmail()\n"
			} else if keyValue == "password" {
				content += "\t@IsStrongPassword()\n"
			}

			if keyValue == "password" || keyValue == "email" {
				content += fmt.Sprintf("\t%s:%s;\n\n", keyName, "string")
			} else {
				content += fmt.Sprintf("\t%s:%s;\n\n", keyName, keyValue)
			}
		}
	}

	if strings.Contains(content, "IsString") {
		classValidatorImports = append(classValidatorImports, "IsString")
	}
	if strings.Contains(content, "number") {
		classValidatorImports = append(classValidatorImports, "IsInt")
	}
	if strings.Contains(content, "IsEmail") {
		classValidatorImports = append(classValidatorImports, "IsEmail")
	}
	if strings.Contains(content, "IsStrongPassword") {
		classValidatorImports = append(classValidatorImports, "IsStrongPassword")
	}

	imports += fmt.Sprintf("\nimport { %s } from 'class-validator';", strings.Join(classValidatorImports, ", "))

	fileName := fmt.Sprintf("%s/dtos/create-%s.dto.ts", name, name)
	fileContent := fmt.Sprintf("%s\n\n%s\n\n%s}", imports, headerContent, content)

	return [2]string{fileName, fileContent}
}
