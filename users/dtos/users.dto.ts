import { CreateUsersDto } from './create-users.dto';
import {ApiProperty} from 'class-validator'
import { PartialType } from '@nestjs/swagger';

export class UpdateUsersDto extends PartialType(CreateUsersDto) {

	@ApiProperty({description: 'Users ID'})
	id: number;

	@ApiProperty({description: 'Users Created Date'})
	createdAt: Date;

	@ApiProperty({description: 'Users Updated Date'})
	UpdatedAt: Date;

	@ApiProperty({description: 'Users Deleted Date', nullable: true})
	deletedAt?: Date;
}