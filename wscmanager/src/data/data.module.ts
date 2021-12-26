import { Module } from '@nestjs/common';
import { DataService } from './data.service';
import { DataController } from './data.controller';
import { ConfigModule } from '@nestjs/config';
import { JwtStrategy } from 'src/middleware/jwt.strategy';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true
    })
  ],
  providers: [DataService, JwtStrategy],
  controllers: [DataController]
})
export class DataModule {}
