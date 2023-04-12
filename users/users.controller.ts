import { Controller, Get, Post, Body, Patch, Param, Delete, Query, Put, UseGuards } from '@nestjs/common';
import { ParseIntPipe } from '@nestjs/common/pipes';
import { UsersDto } from './dto/users.dto';
import {CreateUsersDto} from './dto/create-users.dto';
import {UpdateUsersDto} from './dto/update-users.dto';
import {UsersService} from './users.service'
import { ParseIntPipe } from '@nestjs/common/pipes';
import { UsersDto } from './dto/users.dto';
import {CreateUsersDto} from './dto/create-users.dto';
import {UpdateUsersDto} from './dto/update-users.dto';
import {UsersService} from './users.service'
import { JwtAuthGuard } from 'src/auth/jwt/jwt-auth.guard';
import { ApiBearerAuth, ApiResponse, ApiTags } from '@nestjs/swagger';


@Controller('users')
@UseGuards(JwtAuthGuard)
@ApiTags('Users')
@ApiBearerAuth()
export class UsersController{
	constructor(
		private readonly usersService: UsersService
	){}


	@ApiResponse({ status: 201, description: 'Users created with success.', type: UsersDto})
	@ApiResponse({status: 400, description: 'Provided invalid data.'})
	@Post()
	create(@Body() createUsersDto: CreateUsersDto) {
		return this.usersService.create(createUsersDto);
	}

	@ApiResponse({status: 200, description: 'Users getted with success.', type: [UsersDto]})
	@Get()
	findAll(@Query() query: {relations?: string, take?: number}){
		let relations = query?.relations?.split(',');
		return this.usersService.findAll(relations, query.take);
	}

	@ApiResponse({ status: 200, description: 'Users getted with success.', type: UsersDto})
	@ApiResponse({status: 400, description: 'Provided invalid id.'})
	@Get(':id')
	findOne(@Param('id', ParseIntPipe) id: number, @Query() query: {relations?:string}){
		let relations = query?.relations?.split(',');
		return this.usersService.findOne(id, relations);
	}

	@ApiResponse({ status: 201, description: 'Users updated with success.', type: UsersDto})
	@ApiResponse({status: 400, description: 'Provided invalid data or id.'})
	@Patch(':id')
	update(@Param('id', ParseIntPipe) id: number, @Body() updateUsersDto: UpdateUsersDto){
		return this.usersService.update(id, updateUsersDto);
	}

	@ApiResponse({ status: 201, description: 'Users deleted with successs.'})
	@ApiResponse({status: 400, description: 'Provided invalid id.'})
	@Delete(':id')
	delete(@Param('id', ParseIntPipe) id: number){
		return this.usersService.delete(id);
	}

	@ApiResponse({status: 201, description: 'Users restored with success.'})
	@ApiResponse({status: 400, description: 'Provided invalid id.'})
	@Put('restore/:id')
	restore(@Param('id', ParseIntPipe) id: number){
		return this.usersService.restore(id)
	}

}