package files

import (
	"fmt"
	"strings"
)

/** Create the Dto File*/
func DtoFile(name, singularName string) [2]string {
	titledSingularName := strings.Title(singularName)
	imports := fmt.Sprintf("import { Create%sDto } from './create-%s.dto';\nimport { PartialType, ApiProperty } from '@nestjs/swagger';", titledSingularName, singularName)

	content := fmt.Sprintf("\nexport class %sDto extends PartialType(Create%sDto) {\n", titledSingularName, titledSingularName)

	content += fmt.Sprintf("\n\t@ApiProperty({description: '%s ID'})\n\tid: number;\n\n\t@ApiProperty({description: '%s Created Date'})\n\tcreatedAt: Date;\n\n\t@ApiProperty({description: '%s Updated Date'})\n\tUpdatedAt: Date;\n\n\t@ApiProperty({description: '%s Deleted Date', nullable: true})\n\tdeletedAt?: Date;", titledSingularName, titledSingularName, titledSingularName, titledSingularName)

	fileName := fmt.Sprintf("%s/dtos/%s.dto.ts", name, singularName)
	fileContent := fmt.Sprintf("%s\n%s\n}", imports, content)

	return [2]string{fileName, fileContent}
}
