import { Module } from '@nestjs/common';
import { AddressesService } from './addresses.service';
import { AddressesController } from './addresses.controller';
import { Addresses } from './entities/addresses.entity';
import { TypeOrmModule } from '@nestjs/typeorm';

@Module({	
	imports: [TypeOrmModule.forFeature([Addresses])],
	controllers: [AddressesController],
	providers: [AddressesService],
	exports: [AddressesService]
})
export class AddressesModule { }