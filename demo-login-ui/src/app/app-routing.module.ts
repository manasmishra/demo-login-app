import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SuccessLoginComponent } from './success-login/success-login.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from "./register/register.component";

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'success', component: SuccessLoginComponent },
  { path: "register", component: RegisterComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' }
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
