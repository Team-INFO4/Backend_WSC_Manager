import {IsString, IsNotEmpty} from 'class-validator';

export class DataDto {
    @IsString()
    @IsNotEmpty()
    service: string;

    @IsString()
    @IsNotEmpty()
    startdate: string;

    @IsString()
    @IsNotEmpty()
    enddate: string
}