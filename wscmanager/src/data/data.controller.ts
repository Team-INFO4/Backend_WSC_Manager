import { Body, Controller, Post, UseGuards } from '@nestjs/common';
import { JwtAuthGuard } from 'src/middleware/jwt.guard';
import { DataService } from './data.service';
import { DataDto } from './dto/data.dto';

@Controller('data')
export class DataController {
    constructor(private readonly dataService:DataService) {}

    @Post('find')
    @UseGuards(JwtAuthGuard)
    findData( @Body() body : DataDto){
        return this.dataService.findNotion(body);
    }
}
