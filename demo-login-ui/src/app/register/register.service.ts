import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from '../models/user'

@Injectable({
  providedIn: 'root'
})
export class RegisterService {
  private loginUrl ="http://localhost:8080/users/register";
  constructor(private http: HttpClient) { 
   
  }
  registerUser(user: User): Observable<any> {
    const headers = { 'content-type': 'application/json'};
    const body=JSON.stringify({
      data: {
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
        password: user.password
      }
    });
    console.log(body)
    return this.http.post(this.loginUrl, body,{'headers':headers})
  }
}
