import { ApiProperty } from '@nestjs/swagger';
import { IsOptional, IsNotEmpty, IsString, IsInt } from 'class-validator';

export class CreateUsersDto {


	@ApiProperty({description: 'Users name'})
	@IsNotEmpty()
	@IsString()
	name:string;

	@ApiProperty({description: 'Users age?'})
	@IsOptional()
	@IsInt()
	age?:number;

}