import { Injectable, Inject, InternalServerErrorException } from '@nestjs/common';
import { CreateAddressesDto } from './dtos/addresses.dto
import { UpdateAddressesDto } from './dtos/addresses.dto'
import { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';
import { Logger } from 'winston';

@Injectable()
export class AddressesService {
	constructor(
		@InjectRepository(Addresses) private readonly addressesRepository: Repository<Addresses>,
		@Inject(WINSTON_MODULE_NEST_PROVIDER) private readonly logger: Logger,
	) { }

	async create(createAddressesDto: CreateAddressesDto) {
		try {
			const addresses = this.addressesRepository.create(createAddressesDto);
			return await this.addressesRepository.save(addresses)
		}
		catch (err) {
			this.logger.error(err.message, 'AddressesService');
			throw new InternalServerErrorException();
		}
	}

	async findAll(updateAddressesDto: UpdateAddressesDto, relations?: string[], take: number) {
		try {
			return await this.addressesRepository.find({ relations, take });
		}
		catch (err) {
			this.logger.error(err.message, 'AddressesService');
			throw new InternalServerErrorException();
		}
	}

	async findOne(id: number, relations?: string[]) {
		try {
			return await this.addressesRepository.findOne({ where: { id }, relations });
		}
		catch (err) {
			this.logger.error(err.message, 'AddressesService');
			throw new InternalServerErrorException();
		}
	}

	async update(id: number, updateAddressesDto: UpdateAddressesDto) {
		try {
			await this.addressesRepository.update(id, updateAddressesDto);
			return this.findOne(id);
		}
		catch (err) {
			this.logger.error(err.message, 'AddressesService');
			throw new InternalServerErrorException();
		}
	}

	async delete(id: number) {
		try {
			return await this.addressesRepository.softDelete(id);
		}
		catch (err) {
			this.logger.error(err.message, 'AddressesService');
			throw new InternalServerErrorException();
		}
	}

	async restore(id: number) {
		try {
			return await this.addressesRepository.restore(id);
		}
		catch (err) {
			this.logger.error(err.message, 'AddressesService');
			throw new InternalServerErrorException();
		}
	}

}