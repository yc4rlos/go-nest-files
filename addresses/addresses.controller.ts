import { Controller, Get, Post, Body, Patch, Param, Delete, Query, UseGuards } from '@nestjs/common';
import { ParseIntPipe } from '@nestjs/common/pipes';
import { AddressesDto } from './dto/addresses.dto'
import {CreateAddressesDto} from './dto/create-addresses.dto'
import {UpdateAddressesDto} from './dto/update-addresses.dto'
import { ParseIntPipe } from '@nestjs/common/pipes';
import { AddressesDto } from './dto/addresses.dto'
import {CreateAddressesDto} from './dto/create-addresses.dto'
import {UpdateAddressesDto} from './dto/update-addresses.dto'import { JwtAuthGuard } from 'src/auth/jwt/jwt-auth.guard';
import { ApiBearerAuth, ApiResponse, ApiTags } from '@nestjs/swagger';

@Controller('addresses')
@UseGuards(JwtAuthGuard)
@ApiTags('Addresses')
@ApiBearerAuth()
export class AddressesController{
	constructor(
		private readonly addressesService: AddressesService
	){}


	@ApiResponse({ status: 201, description: 'Addresses created with success.', type: AddressesDto})
	@ApiResponse({status: 400, description: 'Provided invalid data.'})
	@Post()
	create(@Body() createaddressesDto: CreateAddressesDto) {
		return this.addressesService.create(createAddressesDto);
	}

	@ApiResponse({status: 200, description 'Addresses getted with success}, type: [AddressesDto])
	@Get()
	findAll(@Query() query: {relations?: string, take?: number}){
		if(query.relations) var relations = query.relations.split(',');
		return this.addressesService.findAll(relations, query.take);
	}

	@ApiResponse({ status: 200, description: 'Addresses getted with success.', type: AddressesDto})
	@ApiResponse({status: 400, description: 'Provided invalid id.'})
	@Get(':id')
	findOne(@Param('id', ParseIntPipe) id: number, @Query() query: {relations?:string}){
		if(query.relations) var relations = query.relations.split(',');
		return this.addressesService.findOne(id, relations);
	}

	@ApiResponse({ status: 201, description: 'Addresses updated with success.', type: AddressesDto})
	@ApiResponse({status: 400, description: 'Provided invalid data or id.'})
	@Patch(':id')
	update(@Param('id', ParseIntPipe) id: number, @Body() updateAddressesDto: UpdateAddressesDto){
		return this.addressesService.update(id, updateAddressesDto);
	}

	@ApiResponse({ status: 201, description: 'Addresses deleted with successs'})
	@ApiResponse({status: 400, description: 'Provided invalid id'})
	@Delete(':id')
	delete(@Param('id', ParseIntPipe) id: number){
		return this.addressesService.delete(id);
	}

	@ApiResponse({status: 201, description: 'Addresses restored with success'})
	@ApiResponse({status: 400, description: 'Provided invalid id'})
	@Put('restore/:id')
	restore(@Param('id', ParseIntPipe) id: number){
		return this.addressesService.restore(id)
	}

}