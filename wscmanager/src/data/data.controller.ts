import { Body, Controller, Post, Header } from '@nestjs/common';
import { DataService } from './data.service';
import { DataDto } from './dto/data.dto';

@Controller('data')
export class DataController {
    constructor(private readonly dataService:DataService) {}

    @Post('find')
    //@Header('Content-Type', 'text/plain; charset=utf-8')
    findData( @Body() body : DataDto){
        return this.dataService.findNotion(body);
    }
}
