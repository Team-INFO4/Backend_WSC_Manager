import { BadRequestException, ConflictException, UnauthorizedException, Injectable, ForbiddenException, HttpStatus } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from 'src/entity/user.entity';
import { Repository } from 'typeorm';
import { SignupDto } from './dto/signup.dto';
import { hash, compare } from 'bcrypt';
import { LoginDto } from './dto/login.dto';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class AuthService {
    constructor (
        @InjectRepository(User)
        private readonly userRepo: Repository<User>,
        private readonly jwtService: JwtService
    ) {}
    async signup(data : SignupDto) {
        const { id, password, key } = data;
        await this.checkspace(id,password);
        if (key !== process.env.SIGNUPKEY) {
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
            const hashedpassword = await hash(password, process.env.HASH_SALT);
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
    async login (data:LoginDto) {
        const { id, password } = data;
        await this.checkspace(id,password);
        const isUser = await this.userRepo.findOne({
            userid: data.id,
        });
        if (!isUser) {
            throw new ForbiddenException({
                statusCode: HttpStatus.FORBIDDEN,
                message: ['Wrong Auth']
            });
        }
        await this.verifyPassword(data.password, isUser.password);
        
        const jwt = await this.jwtService.signAsync({
            id: data.id
        });
        return {
            message: "Login success",
            accessToken: jwt
        }
    }

    private async checkspace(id:string, password:string) {
        if (id.indexOf(' ') !== -1 || password.indexOf(' ') !== -1 ) {
            throw new BadRequestException();
        }
    }
    private async verifyPassword(plainpassword:string, hashpassword:string){
        const isMatch = await compare(plainpassword, hashpassword);
        if(!isMatch){
            throw new ForbiddenException();
        }
    }
}
