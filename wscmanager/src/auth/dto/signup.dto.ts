import {IsString, IsNotEmpty} from 'class-validator';

export class SignupDto {
    @IsString()
    @IsNotEmpty()
    id: string;

    @IsString()
    @IsNotEmpty()
    password: string;

    @IsString()
    @IsNotEmpty()
    key: string;
}