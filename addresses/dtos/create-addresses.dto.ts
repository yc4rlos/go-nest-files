import { ApiProperty } from '@nestjs/swagger';
import { IsOptional, IsNotEmpty, IsString } from 'class-validator';

export class createAddressesDto {


	@ApiProperty({description: 'Addresses rua'})
	@IsNotEmpty()
	@IsString()
	rua:string;

	@ApiProperty({description: 'Addresses uf'})
	@IsNotEmpty()
	@IsString()
	uf:string;

	@ApiProperty({description: 'Addresses logradouro?'})
	@IsOptional()
	logradouro?:int;

	@ApiProperty({description: 'Addresses cep'})
	@IsNotEmpty()
	@IsString()
	cep:string;

	@ApiProperty({description: 'Addresses cidade'})
	@IsNotEmpty()
	@IsString()
	cidade:string;

}