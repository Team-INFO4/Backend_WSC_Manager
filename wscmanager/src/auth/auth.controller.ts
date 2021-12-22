import { Body, Controller, HttpCode, HttpStatus, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { LoginDto } from './dto/login.dto';
import { SignupDto } from './dto/signup.dto';

@Controller('auth')
export class AuthController {
    constructor(private readonly authService:AuthService){}

    @Post('signup')
    @HttpCode(HttpStatus.CREATED)
    signup( @Body() body : SignupDto){
        return this.authService.signup(body);
    }

    @Post('login')
    login( @Body() body : LoginDto){
        return this.authService.login(body);
    }
}
