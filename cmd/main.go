package main

import (
	"fmt"
	"os"
	"strings"
)

/** Defines if it will have authentication*/
var auth bool

/** Defines if it will have documentation*/
var documentation bool

/** Defines if have error log config (With Nest Winston)*/
var logger bool

/** Resource Name*/
var name string

/** Titled Resource Name*/
var titledName string

/** Properties List*/
var properties = make([]string, 2, 15)

/** Folder Path */
var folderPath string

func main() {

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		if arg[0] == '-' {
			switch arg {
			case "-af", "-all-flags":
				auth = true
				documentation = true
				logger = true
			case "-a", "-auth":
				auth = true
			case "-d", "-docs":
				documentation = true
			case "-l", "-log":
				logger = true
			case "-h", "-help":
				help()
				return
			default:
				fmt.Printf("Command %s invalid, use -h for more information.\n", arg)
			}
		} else {
			if name == "" {
				name = arg
				titledName = strings.Title(arg)
				continue
			}

			if i == len(os.Args)-1 {

				if strings.ContainsRune(arg, '/') {
					folderPath = arg
				} else {
					fmt.Println("The Path cannot be empty.")
					os.Exit(0)
				}
			}

			if strings.ContainsRune(arg, ':') {
				properties = append(properties, arg)
			}

		}
	}

	createFolders()

	controllerFile()
	serviceFile()
	moduleFile()
	createEntityDtoFile()
	updateEntityDtoFile()
	entityFile()
	if documentation {
		entityDtoFile()
	}

	fmt.Println("Files generated!")
}

/** Create the Folder and files */
func createFolders() {

	// Create the main Folder
	os.Mkdir(folderPath+name, 0777)

	// Create the dto folder and files
	os.Mkdir(folderPath+name+"/dtos", 0777)

	// Create the Entity Folder and File
	os.Mkdir(folderPath+name+"/entities", 0777)
}

/** Print in console the Help Options*/
func help() {
	fmt.Println(`Command List:
	All the flags: -af | -all-flags;
	Flag to add the Authentication: -a | -auth;
	Flag to add the Documentation: -d  | -docs;
	Flag to add the Logger: -l | -log
	
	Avaible Types: number | string | email | password
	Optional Declaration "?:"
	Required Declaration ":"
	The directory path should be the last argument to the command
	`)
}

/** Create the Controller file*/
func controllerFile() {
	// TODO: change this code latter to set dynamically the @nestjs/common imports
	nestCommonImports := []string{"Controller", "Get", "Post", "Body", "Patch", "Param", "Delete", "Query", "Put"}
	imports := fmt.Sprintf("\nimport { ParseIntPipe } from '@nestjs/common/pipes';\nimport { %sDto } from './dtos/%s.dto';\nimport {Create%sDto} from './dtos/create-%s.dto';\nimport {Update%sDto} from './dtos/update-%s.dto';\nimport {%sService} from './%s.service'", titledName, name, titledName, name, titledName, name, titledName, name)
	authDocumentation := ""
	headerDecorators := fmt.Sprintf("@Controller('%s')\n", name)
	content := fmt.Sprintf("export class %sController{\n\tconstructor(\n\t\tprivate readonly %sService: %sService\n\t){}\n", titledName, name, titledName)
	routes := [6]string{}

	// * Route for Create
	routes[0] = fmt.Sprintf("\t@Post()\n\tcreate(@Body() create%sDto: Create%sDto) {\n\t\treturn this.%sService.create(create%sDto);\n\t}", titledName, titledName, name, titledName)

	// * Route for FindAll
	routes[1] = fmt.Sprintf("\t@Get()\n\tfindAll(@Query() query: {relations?: string, take?: number}){\n\t\tlet relations = query?.relations?.split(',');\n\t\treturn this.%sService.findAll(relations, query.take);\n\t}", name)

	// * Rotue for FindOne
	routes[2] = fmt.Sprintf("\t@Get(':id')\n\tfindOne(@Param('id', ParseIntPipe) id: number, @Query() query: {relations?:string}){\n\t\tlet relations = query?.relations?.split(',');\n\t\treturn this.%sService.findOne(id, relations);\n\t}", name)

	// * Route for Update
	routes[3] = fmt.Sprintf("\t@Patch(':id')\n\tupdate(@Param('id', ParseIntPipe) id: number, @Body() update%sDto: Update%sDto){\n\t\treturn this.%sService.update(id, update%sDto);\n\t}", titledName, titledName, name, titledName)

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
		routes[0] = fmt.Sprintf("\t@ApiResponse({ status: 201, description: '%s created with success.', type: %sDto})\n\t@ApiResponse({status: 400, description: 'Provided invalid data.'})\n%s", titledName, titledName, routes[0])

		// * Route for FindAll
		routes[1] = fmt.Sprintf("\t@ApiResponse({status: 200, description: '%s getted with success.', type: [%sDto]})\n%s", titledName, titledName, routes[1])

		// * Rotue for FindOne
		routes[2] = fmt.Sprintf("\t@ApiResponse({ status: 200, description: '%s getted with success.', type: %sDto})\n\t@ApiResponse({status: 400, description: 'Provided invalid id.'})\n%s", titledName, titledName, routes[2])

		// * Route for Update
		routes[3] = fmt.Sprintf("\t@ApiResponse({ status: 201, description: '%s updated with success.', type: %sDto})\n\t@ApiResponse({status: 400, description: 'Provided invalid data or id.'})\n%s", titledName, titledName, routes[3])

		// * Route for Delete
		routes[4] = fmt.Sprintf("\t@ApiResponse({ status: 201, description: '%s deleted with successs.'})\n\t@ApiResponse({status: 400, description: 'Provided invalid id.'})\n%s", titledName, routes[4])

		// * Route for Restore
		routes[5] = fmt.Sprintf("\t@ApiResponse({status: 201, description: '%s restored with success.'})\n\t@ApiResponse({status: 400, description: 'Provided invalid id.'})\n%s", titledName, routes[5])
	}

	for _, route := range routes {
		content += "\n\n" + route
	}

	imports = fmt.Sprintf("import { %s } from '@nestjs/common';%s\n", strings.Join(nestCommonImports, ", "), imports)

	fileName := fmt.Sprintf("%s/%s.controller.ts", name, name)
	fileContent := fmt.Sprintf("%s\n%s%s\n}", imports, headerDecorators, content)
	createFile(fileName, fileContent)
}

/** Create the Service file */
func serviceFile() {

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

	createFile(fileName, fileContent)
}

/** Create Dto file*/
func createEntityDtoFile() {

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
				content += fmt.Sprintf("\t@ApiProperty({description: '%s %s'})\n", titledName, keyName)
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

	createFile(fileName, fileContent)
}

/** Create the UpdateEntity file*/
func updateEntityDtoFile() {
	imports := fmt.Sprintf("import { Create%sDto } from './create-%s.dto';\nimport { PartialType } from '@nestjs/swagger';\n", titledName, name)

	content := fmt.Sprintf("export class Update%sDto extends PartialType(Create%sDto) {}", titledName, titledName)

	fileName := fmt.Sprintf("%s/dtos/update-%s.dto.ts", name, name)
	fileContent := fmt.Sprintf("%s\n%s\n", imports, content)
	createFile(fileName, fileContent)
}

/** Create the entity Dto File*/
func entityDtoFile() {
	imports := fmt.Sprintf("import { Create%sDto } from './create-%s.dto';\nimport { PartialType, ApiProperty } from '@nestjs/swagger';", titledName, name)

	content := fmt.Sprintf("\nexport class %sDto extends PartialType(Create%sDto) {\n", titledName, titledName)

	content += fmt.Sprintf("\n\t@ApiProperty({description: '%s ID'})\n\tid: number;\n\n\t@ApiProperty({description: '%s Created Date'})\n\tcreatedAt: Date;\n\n\t@ApiProperty({description: '%s Updated Date'})\n\tUpdatedAt: Date;\n\n\t@ApiProperty({description: '%s Deleted Date', nullable: true})\n\tdeletedAt?: Date;", titledName, titledName, titledName, titledName)

	fileContent := fmt.Sprintf("%s\n%s\n}", imports, content)
	fileName := fmt.Sprintf("%s/dtos/%s.dto.ts", name, name)

	createFile(fileName, fileContent)
}

/** Create the Entity file */
func entityFile() {
	imports := "import { PrimaryGeneratedColumn, CreateDateColumn, UpdateDateColumn, DeleteDateColumn, Column, Entity } from 'typeorm';\n"
	content := fmt.Sprintf("@Entity()\nexport class %s{\n\n\t@PrimaryGeneratedColumn()\n\tid:number;\n\n", titledName)

	for _, property := range properties {
		if property != "" {

			colonIndex := strings.IndexRune(property, ':')
			optional := strings.ContainsRune(property, '?')
			keyName := property[:colonIndex]
			keyValue := property[colonIndex+1:]

			if optional {
				content += "\t@Column({nullable: true})\n"
			} else {
				content += "\t@Column()\n"
			}
			if keyValue == "password" || keyValue == "email" {
				content += fmt.Sprintf("\t%s:%s;\n\n", keyName, "string")
			} else {
				content += fmt.Sprintf("\t%s:%s;\n\n", keyName, keyValue)
			}
		}
	}

	content += "\t@CreateDateColumn()\n\tcreatedAt: Date;\n\n\t@UpdateDateColumn()\n\tupdatedAt: Date;\n\n\t@DeleteDateColumn()\n\tdeletedAt: Date;"

	fileContent := fmt.Sprintf("%s\n%s\n}", imports, content)
	fileName := fmt.Sprintf("%s/entities/%s.entity.ts", name, name)

	createFile(fileName, fileContent)
}

/** Create the Module file*/
func moduleFile() {
	imports := fmt.Sprintf("import { Module } from '@nestjs/common';\nimport { %sService } from './%s.service';\nimport { %sController } from './%s.controller';\nimport { %s } from './entities/%s.entity';\nimport { TypeOrmModule } from '@nestjs/typeorm';\n\n", titledName, name, titledName, name, titledName, name)

	content := fmt.Sprintf("@Module({\t\n\timports: [TypeOrmModule.forFeature([%s])],\n\tcontrollers: [%sController],\n\tproviders: [%sService],\n\texports: [%sService]\n})\nexport class %sModule { }", titledName, titledName, titledName, titledName, titledName)

	fileContent := imports + content
	fileName := fmt.Sprintf("%s/%s.module.ts", name, name)

	createFile(fileName, fileContent)
}

/** Create the Files */
func createFile(name, content string) {

	file, err := os.Create(folderPath + name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileBytes := []byte(content)
	file.Write(fileBytes)

	file.Close()
}
