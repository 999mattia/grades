import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { AuthGuard } from '../shared/auth.guard';
import { CreateModuleComponent } from './create-module/create-module.component';
import { CreateGradeComponent } from './create-grade/create-grade.component';

const routes: Routes = [
	{ path: '', component: HomeComponent, canActivate: [AuthGuard] },
	{ path: 'login', component: LoginComponent },
	{
		path: 'module/create',
		component: CreateModuleComponent,
		canActivate: [AuthGuard],
	},
	{
		path: 'module/:id/create',
		component: CreateGradeComponent,
		canActivate: [AuthGuard],
	},
];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule],
})
export class AppRoutingModule {}
