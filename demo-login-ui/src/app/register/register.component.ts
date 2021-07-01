import { Component, OnInit } from '@angular/core';
import { RegisterService } from "./register.service"
import { Router, ActivatedRoute } from '@angular/router';
import {User} from '../models/user';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  user:User={email:'',password:'', firstname: "", lastname: ""};
  successData:any;
  returnUrl = "/login";
  message = "";
  messagesuccess = "";

  constructor(private registerService: RegisterService,
    private route: ActivatedRoute,
    private router: Router) { }

  ngOnInit(): void {
  }

  registerMethod() {
    this.message = "";
    this.messagesuccess = "";
    if (!this.user.email || !this.user.password || !this.user.firstname || !this.user.lastname) {
      this.message = "Please input all the required fields";
      return
    }
    this.registerService.registerUser (this.user)
      .subscribe(data => {
        console.log(data);
        this.successData = data;
        this.router.navigate([this.returnUrl]);
        this.messagesuccess = "You have sucessfully registered please login now!!!"
        // this.refreshPeople();
      },
      err => {
        console.log('HTTP Error', err)
        this.message = err.message;

      },
      () => console.log('HTTP request completed.'))      
  }

}
