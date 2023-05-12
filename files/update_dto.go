package files

import (
	"fmt"
)

/** Create the UpdateEntity file*/
func UpdateDtoFile(titledName, name string) [2]string {
	imports := fmt.Sprintf("import { Create%sDto } from './create-%s.dto';\nimport { PartialType } from '@nestjs/swagger';\n", titledName, name)

	content := fmt.Sprintf("export class Update%sDto extends PartialType(Create%sDto) {}", titledName, titledName)

	fileName := fmt.Sprintf("%s/dtos/update-%s.dto.ts", name, name)
	fileContent := fmt.Sprintf("%s\n%s\n", imports, content)
	return [2]string{fileName, fileContent}
}
