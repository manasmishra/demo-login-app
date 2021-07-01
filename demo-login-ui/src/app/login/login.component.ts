import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { LoginServiceService } from './login-service.service';
import {User} from '../models/user';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  
  user:User={email:'',password:'', firstname: "", lastname: ""};
  email = "";
  successData:any;
  returnUrl = "/success";
  message = ""
  constructor(
    private loginService: LoginServiceService,
    private cookieService: CookieService,
    private route: ActivatedRoute,
    private router: Router) {
    
   }

  ngOnInit(): void {
    // this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/success';
  }
  loginMethod() {
    this.message = "";
    if (!this.user.email || !this.user.password ) {
      this.message = "Please input all the required fields";
      return
    }
    this.loginService.getUser(this.user)
      .subscribe(data => {
        // console.log(data);
        this.successData = data;
        this.cookieService.set('userToken',this.successData.data.token);
        // console.log('cookie', this.cookieService.get('userToken'));
        this.router.navigate([this.returnUrl]);
        // this.refreshPeople();
      },
      err => {
        // console.log('HTTP Error', err)
        this.message = err.message;
      },
      () => console.log('HTTP request completed.'))      
  }

}
