import { Injectable, Inject, InternalServerErrorException} from '@nestjs/common';
import { CreateUsersDto } from './dtos/users.dto';
import { UpdateUsersDto } from './dtos/users.dto';
import { Users } from './entities/users.entity'
import { Repository } from 'typeorm';
import { InjectRepository } from '@nestjs/typeorm';
import { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';
import { Logger } from 'winston';

@Injectable()
export class UsersService {
	constructor(
		@InjectRepository(Users) private readonly usersRepository: Repository<Users>,
		@Inject(WINSTON_MODULE_NEST_PROVIDER) private readonly logger: Logger,
	){}

	async create(createUsersDto: CreateUsersDto){
		try{
			const users = this.usersRepository.create(createUsersDto);
			return await this.usersRepository.save(users)
	}
		catch(err){
			this.logger.error(err.message, 'UsersService');
			throw new InternalServerErrorException();
		}
	}

	async findAll(relations?: string[], take?: number){
		try{
			return await this.usersRepository.find({relations, take});
	}
		catch(err){
			this.logger.error(err.message, 'UsersService');
			throw new InternalServerErrorException();
		}
	}

	async findOne(id: number, relations?: string[]){
		try{
			return await this.usersRepository.findOne({where:{id},relations});
	}
		catch(err){
			this.logger.error(err.message, 'UsersService');
			throw new InternalServerErrorException();
		}
	}

	async update(id: number, updateUsersDto: UpdateUsersDto){
		try{
			await this.usersRepository.update(id, updateUsersDto);
			return this.findOne(id);
	}
		catch(err){
			this.logger.error(err.message, 'UsersService');
			throw new InternalServerErrorException();
		}
	}

	async delete(id: number){
		try{
			return await this.usersRepository.softDelete(id);
		}
		catch(err){
			this.logger.error(err.message, 'UsersService');
			throw new InternalServerErrorException();
		}
	}

	async restore(id: number){
		try{
			return await this.usersRepository.restore(id);
		}
		catch(err){
			this.logger.error(err.message, 'UsersService');
			throw new InternalServerErrorException();
		}
	}

}