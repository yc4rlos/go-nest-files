package files

import (
	"fmt"
	"strings"
)

/** Create the UpdateEntity file*/
func UpdateDtoFile(name, singularName string) [2]string {
	titledSingularName := strings.Title(singularName)
	imports := fmt.Sprintf("import { Create%sDto } from './create-%s.dto';\nimport { PartialType } from '@nestjs/swagger';\n", titledSingularName, singularName)

	content := fmt.Sprintf("export class Update%sDto extends PartialType(Create%sDto) {}", titledSingularName, titledSingularName)

	fileName := fmt.Sprintf("%s/dtos/update-%s.dto.ts", name, singularName)
	fileContent := fmt.Sprintf("%s\n%s\n", imports, content)
	return [2]string{fileName, fileContent}
}
