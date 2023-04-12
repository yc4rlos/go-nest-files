import { PrimaryGeneratedColumn, CreateDateColumn, UpdateDateColumn, DeleteDateColumn, Column, Entity } from 'typeorm';

@Entity()
export class Addresses{

	@PrimaryGeneratedColumn()
	id:number;

	@Column()
	rua:string;

	@Column()
	uf:string;

	@Column({nullable: true})
	logradouro?:int;

	@Column()
	cep:string;

	@Column()
	cidade:string;

	@CreateDateColumn()
	createdAt: Date;

	@UpdateDateColumn()
	updatedAt: Date;

	@DeleteDateColumn()
	deletedAt: Date;
}