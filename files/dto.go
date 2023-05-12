package files

import (
	"fmt"
)

/** Create the Dto File*/
func DtoFile(titledName, name string) [2]string {
	imports := fmt.Sprintf("import { Create%sDto } from './create-%s.dto';\nimport { PartialType, ApiProperty } from '@nestjs/swagger';", titledName, name)

	content := fmt.Sprintf("\nexport class %sDto extends PartialType(Create%sDto) {\n", titledName, titledName)

	content += fmt.Sprintf("\n\t@ApiProperty({description: '%s ID'})\n\tid: number;\n\n\t@ApiProperty({description: '%s Created Date'})\n\tcreatedAt: Date;\n\n\t@ApiProperty({description: '%s Updated Date'})\n\tUpdatedAt: Date;\n\n\t@ApiProperty({description: '%s Deleted Date', nullable: true})\n\tdeletedAt?: Date;", titledName, titledName, titledName, titledName)

	fileContent := fmt.Sprintf("%s\n%s\n}", imports, content)
	fileName := fmt.Sprintf("%s/dtos/%s.dto.ts", name, name)

	return [2]string{fileName, fileContent}
}
