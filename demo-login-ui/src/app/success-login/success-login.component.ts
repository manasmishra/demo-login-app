import { Component, OnInit } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-success-login',
  templateUrl: './success-login.component.html',
  styleUrls: ['./success-login.component.css']
})
export class SuccessLoginComponent implements OnInit {

  constructor(private cookieService: CookieService) { }

  ngOnInit(): void {
  }
  logoutMethod(){
    console.log('cookie', this.cookieService.get('userToken'));
    this.cookieService.delete('userToken');
  }

}
