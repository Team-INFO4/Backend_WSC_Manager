import { BadRequestException, ConflictException, UnauthorizedException, Injectable, ForbiddenException, HttpStatus } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from 'src/entity/user.entity';
import { Repository } from 'typeorm';
import { SignupDto } from './dto/signup.dto';
import { hash, compare } from 'bcrypt';

@Injectable()
export class AuthService {
    constructor (
        @InjectRepository(User)
        private readonly userRepo: Repository<User>
    ) {}
    async signup(data : SignupDto) {
        const { id, password, key } = data;
        console.log(id,password);
        if (id.indexOf(' ') !== -1 || password.indexOf(' ') !== -1 || key !== process.env.SIGNUPKEY) {
            throw new BadRequestException();
        }
        const isExist = await this.userRepo.findOne({
            userid: data.id
        });
        if (isExist) {
            throw new ForbiddenException({
                statusCode: HttpStatus.FORBIDDEN,
                message: ['Already registered user.']
            });
        }
        try {
            const hashedpassword = await hash(password, 10);
            await this.userRepo.save({
                userid: id,
                password: hashedpassword
            })
        } catch (error) {
            return {
                ...error
            };
        }
        return {
            statusCode: HttpStatus.CREATED
        };
        
    }
}
