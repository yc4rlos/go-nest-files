package files

import (
	"fmt"
)

/** Create the Module file*/
func ModuleFile(titledName, name string) [2]string {
	imports := fmt.Sprintf("import { Module } from '@nestjs/common';\nimport { %sService } from './%s.service';\nimport { %sController } from './%s.controller';\nimport { %s } from './entities/%s.entity';\nimport { TypeOrmModule } from '@nestjs/typeorm';\n\n", titledName, name, titledName, name, titledName, name)

	content := fmt.Sprintf("@Module({\t\n\timports: [TypeOrmModule.forFeature([%s])],\n\tcontrollers: [%sController],\n\tproviders: [%sService],\n\texports: [%sService]\n})\nexport class %sModule { }", titledName, titledName, titledName, titledName, titledName)

	fileContent := imports + content
	fileName := fmt.Sprintf("%s/%s.module.ts", name, name)

	return [2]string{fileName, fileContent}
}
