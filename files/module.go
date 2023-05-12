package files

import (
	"fmt"
	"strings"
)

/** Create the Module file*/
func ModuleFile(name, singularName string) [2]string {
	titledName := strings.Title(name)
	titledSingularName := strings.Title(singularName)
	imports := fmt.Sprintf("import { Module } from '@nestjs/common';\nimport { %sService } from './%s.service';\nimport { %sController } from './%s.controller';\nimport { %s } from './entities/%s.entity';\nimport { TypeOrmModule } from '@nestjs/typeorm';\n\n", titledName, name, titledName, name, titledSingularName, singularName)

	content := fmt.Sprintf("@Module({\t\n\timports: [TypeOrmModule.forFeature([%s])],\n\tcontrollers: [%sController],\n\tproviders: [%sService],\n\texports: [%sService]\n})\nexport class %sModule { }", titledSingularName, titledName, titledName, titledName, titledName)

	fileContent := imports + content
	fileName := fmt.Sprintf("%s/%s.module.ts", name, name)

	return [2]string{fileName, fileContent}
}
