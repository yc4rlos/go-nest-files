import { PrimaryGeneratedColumn, CreateDateColumn, UpdateDateColumn, DeleteDateColumn, Column, Entity } from 'typeorm';

@Entity()
export class Users{

	@PrimaryGeneratedColumn()
	id:number;

	@Column()
	name:string;

	@Column({nullable: true})
	age?:number;

	@CreateDateColumn()
	createdAt: Date;

	@UpdateDateColumn()
	updatedAt: Date;

	@DeleteDateColumn()
	deletedAt: Date;
}