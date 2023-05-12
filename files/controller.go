package files

import (
	"fmt"
	"strings"
)

/** Create the Controller file*/
func ControllerFile(name, singularName string, auth, documentation, logger bool) [2]string {

	titledName := strings.Title(name)
	titledSingularName := strings.Title(singularName)

	nestCommonImports := []string{"Controller", "Get", "Post", "Body", "Patch", "Param", "Delete", "Query", "Put"}
	imports := fmt.Sprintf("\nimport { ParseIntPipe } from '@nestjs/common/pipes';\nimport { %sDto } from './dtos/%s.dto';\nimport {Create%sDto} from './dtos/create-%s.dto';\nimport {Update%sDto} from './dtos/update-%s.dto';\nimport {%sService} from './%s.service'", titledSingularName, singularName, titledSingularName, singularName, titledSingularName, singularName, titledName, name)
	authDocumentation := ""
	headerDecorators := fmt.Sprintf("@Controller('%s')\n", name)
	content := fmt.Sprintf("export class %sController{\n\tconstructor(\n\t\tprivate readonly %sService: %sService\n\t){}\n", titledName, name, titledName)
	routes := [6]string{}

	// * Route for Create
	routes[0] = fmt.Sprintf("\t@Post()\n\tcreate(@Body() create%sDto: Create%sDto) {\n\t\treturn this.%sService.create(create%sDto);\n\t}", titledSingularName, titledSingularName, name, titledSingularName)

	// * Route for FindAll
	routes[1] = fmt.Sprintf("\t@Get()\n\tfindAll(@Query() query: {relations?: string, take?: number}){\n\t\tlet relations = query?.relations?.split(',');\n\t\treturn this.%sService.findAll(relations, query.take);\n\t}", name)

	// * Rotue for FindOne
	routes[2] = fmt.Sprintf("\t@Get(':id')\n\tfindOne(@Param('id', ParseIntPipe) id: number, @Query() query: {relations?:string}){\n\t\tlet relations = query?.relations?.split(',');\n\t\treturn this.%sService.findOne(id, relations);\n\t}", name)

	// * Route for Update
	routes[3] = fmt.Sprintf("\t@Patch(':id')\n\tupdate(@Param('id', ParseIntPipe) id: number, @Body() update%sDto: Update%sDto){\n\t\treturn this.%sService.update(id, update%sDto);\n\t}", titledSingularName, titledSingularName, name, titledSingularName)

	// * Route for Delete
	routes[4] = fmt.Sprintf("\t@Delete(':id')\n\tdelete(@Param('id', ParseIntPipe) id: number){\n\t\treturn this.%sService.delete(id);\n\t}", name)

	// * Route for Restore
	routes[5] = fmt.Sprintf("\t@Put('restore/:id')\n\trestore(@Param('id', ParseIntPipe) id: number){\n\t\treturn this.%sService.restore(id)\n\t}\n", name)

	if auth {
		nestCommonImports = append(nestCommonImports, "UseGuards")
		imports = fmt.Sprintf("%s\nimport { JwtAuthGuard } from 'src/auth/jwt/jwt-auth.guard';\n", imports)
		headerDecorators = fmt.Sprintf("%s@UseGuards(JwtAuthGuard)\n", headerDecorators)
		authDocumentation = "ApiBearerAuth, "
	}

	if documentation {
		imports = fmt.Sprintf("%simport { %sApiResponse, ApiTags } from '@nestjs/swagger';\n", imports, authDocumentation)
		headerDecorators = fmt.Sprintf("%s@ApiTags('%s')\n", headerDecorators, titledName)
		if auth {
			headerDecorators = fmt.Sprintf("%s@ApiBearerAuth()\n", headerDecorators)
		}

		// * Route for Create
		routes[0] = fmt.Sprintf("\t@ApiResponse({ status: 201, description: '%s created with success.', type: %sDto})\n\t@ApiResponse({status: 400, description: 'Provided invalid data.'})\n%s", titledSingularName, titledSingularName, routes[0])

		// * Route for FindAll
		routes[1] = fmt.Sprintf("\t@ApiResponse({status: 200, description: '%s getted with success.', type: [%sDto]})\n%s", titledName, titledSingularName, routes[1])

		// * Rotue for FindOne
		routes[2] = fmt.Sprintf("\t@ApiResponse({ status: 200, description: '%s getted with success.', type: %sDto})\n\t@ApiResponse({status: 400, description: 'Provided invalid id.'})\n%s", titledSingularName, titledSingularName, routes[2])

		// * Route for Update
		routes[3] = fmt.Sprintf("\t@ApiResponse({ status: 201, description: '%s updated with success.', type: %sDto})\n\t@ApiResponse({status: 400, description: 'Provided invalid data or id.'})\n%s", titledSingularName, titledSingularName, routes[3])

		// * Route for Delete
		routes[4] = fmt.Sprintf("\t@ApiResponse({ status: 201, description: '%s deleted with successs.'})\n\t@ApiResponse({status: 400, description: 'Provided invalid id.'})\n%s", titledSingularName, routes[4])

		// * Route for Restore
		routes[5] = fmt.Sprintf("\t@ApiResponse({status: 201, description: '%s restored with success.'})\n\t@ApiResponse({status: 400, description: 'Provided invalid id.'})\n%s", titledSingularName, routes[5])
	}

	for _, route := range routes {
		content += "\n\n" + route
	}

	imports = fmt.Sprintf("import { %s } from '@nestjs/common';%s\n", strings.Join(nestCommonImports, ", "), imports)

	fileName := fmt.Sprintf("%s/%s.controller.ts", name, name)
	fileContent := fmt.Sprintf("%s\n%s%s\n}", imports, headerDecorators, content)

	return [2]string{fileName, fileContent}
}
