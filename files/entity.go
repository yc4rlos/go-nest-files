package files

import (
	"fmt"
	"strings"
)

/** Create the Entity file */
func EntityFile(name, singularName string, properties []string) [2]string {
	titledSingularName := strings.Title(singularName)
	imports := "import { PrimaryGeneratedColumn, CreateDateColumn, UpdateDateColumn, DeleteDateColumn, Column, Entity } from 'typeorm';\n"
	content := fmt.Sprintf("@Entity()\nexport class %s{\n\n\t@PrimaryGeneratedColumn()\n\tid:number;\n\n", titledSingularName)

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
	fileName := fmt.Sprintf("%s/entities/%s.entity.ts", name, singularName)

	return [2]string{fileName, fileContent}
}
