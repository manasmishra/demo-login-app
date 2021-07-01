import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from '../models/user'

@Injectable({
  providedIn: 'root'
})
export class LoginServiceService {
  private loginUrl ="http://localhost:8080/users/login";
  constructor(private http: HttpClient) { 
   
  }
  getUser(user: User): Observable<any> {
    const headers = { 'content-type': 'application/json'};
    const body=JSON.stringify({
      data: {
        email: user.email,
        password: user.password
      }
    });
    console.log(body)
    return this.http.post(this.loginUrl, body,{'headers':headers})
  }
}
