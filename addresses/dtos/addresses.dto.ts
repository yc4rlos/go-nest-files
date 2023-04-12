import { CreateAddressesDto } from './create-addresses.dto';
import {ApiProperty} from 'class-validator'

export class UpdateAddressesDto extends PartialType(CreateAddressesDto) {

	@ApiProperty({description: 'Addresses ID'})
	id: number;

	@ApiProperty({description: 'Addresses Created Date'})
	createdAt: Date;

	@ApiProperty({description: 'Addresses Updated Date'})
	UpdatedAt: Date;

	@ApiProperty({description: 'Addresses Deleted Date', nullable: true})
	deletedAt?: Date;
}