import { Body, Controller, HttpCode, HttpStatus, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { SignupDto } from './dto/signup.dto';

@Controller('auth')
export class AuthController {
    constructor(private readonly authService:AuthService){}

    @Post('signup')
    @HttpCode(HttpStatus.CREATED)
    signup( @Body() body : SignupDto){
        return this.authService.signup(body);
    }
}
