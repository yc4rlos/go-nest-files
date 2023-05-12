package files

import (
	"fmt"
	"strings"
)

/** Create the Service file */
func ServiceFile(titledName, name string, auth, documentation, logger bool) [2]string {

	nestCommonImports := []string{"Injectable"}
	imports := fmt.Sprintf("import { Create%sDto } from './dtos/create-%s.dto';\nimport { Update%sDto } from './dtos/update-%s.dto';\nimport { %s } from './entities/%s.entity'\nimport { Repository } from 'typeorm';\nimport { InjectRepository } from '@nestjs/typeorm';\n", titledName, name, titledName, name, titledName, name)
	headerDecorators := "@Injectable()\n"
	constructorContent := fmt.Sprintf("@InjectRepository(%s) private readonly %sRepository: Repository<%s>,", titledName, name, titledName)

	methods := [6]string{}

	// * Method to create
	methods[0] = fmt.Sprintf("\tasync create(create%sDto: Create%sDto){\n\t\t\tconst %s = this.%sRepository.create(create%sDto);\n\t\t\treturn await this.%sRepository.save(%s)\n\t}", titledName, titledName, name, name, titledName, name, name)

	// * Method to Find All
	methods[1] = fmt.Sprintf("\tasync findAll(relations?: string[], take?: number){\n\t\t\treturn await this.%sRepository.find({relations, take});\n\t}", name)

	// * Method to Find One
	methods[2] = fmt.Sprintf("\tasync findOne(id: number, relations?: string[]){\n\t\t\treturn await this.%sRepository.findOne({where:{id},relations});\n\t}", name)

	// * Method to Update
	methods[3] = fmt.Sprintf("\tasync update(id: number, update%sDto: Update%sDto){\n\t\t\tawait this.%sRepository.update(id, update%sDto);\n\t\t\treturn this.findOne(id);\n\t}", titledName, titledName, name, titledName)

	// * Method to Delete
	methods[4] = fmt.Sprintf("\tasync delete(id: number){\n\t\t\treturn await this.%sRepository.softDelete(id);\n\t\t}", name)

	// * Method to Restore
	methods[5] = fmt.Sprintf("\tasync restore(id: number){\n\t\t\treturn await this.%sRepository.restore(id);\n\t\t}", name)

	// Is logger be True the lines for winston logger will be inserted
	if logger {
		nestCommonImports = append(nestCommonImports, "Inject", "InternalServerErrorException")
		imports = fmt.Sprintf("%simport { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';\nimport { Logger } from 'winston';\n", imports)
		constructorContent = fmt.Sprintf("%s\n\t\t@Inject(WINSTON_MODULE_NEST_PROVIDER) private readonly logger: Logger,", constructorContent)

		for i, method := range methods {
			text := strings.IndexRune(method, '{')
			if text != -1 {
				methods[i] = fmt.Sprintf("%s\t\ttry{\n%s\n\t\tcatch(err){\n\t\t\tthis.logger.error(err.message, '%sService');\n\t\t\tthrow new InternalServerErrorException();\n\t\t}\n\t}\n", method[:text+2], method[text+2:], titledName)
			}
		}
	}

	content := fmt.Sprintf("export class %sService {\n\tconstructor(\n\t\t%s\n\t){}\n", titledName, constructorContent)

	for _, method := range methods {
		content += "\n" + method
	}

	imports = fmt.Sprintf("import { %s} from '@nestjs/common';\n%s", strings.Join(nestCommonImports, ", "), imports)

	fileName := fmt.Sprintf("%s/%s.service.ts", name, name)
	fileContent := fmt.Sprintf("%s\n%s%s\n}", imports, headerDecorators, content)

	return [2]string{fileName, fileContent}
}
