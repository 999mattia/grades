import { Component, HostListener } from '@angular/core';
import { AuthService } from '../shared/auth.service';
import { interval } from 'rxjs';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
})
export class AppComponent {
	constructor(private authService: AuthService) {}

	title = 'Grades';
	username = '';

	handleLogout() {
		this.authService.logOut();
		this.username = '';
	}

	ngOnInit(): void {
		this.authService.getCurrentUsername()?.subscribe((res: any) => {
			this.username = res.username;
		});

		interval(1000).subscribe(() => {
			if (!this.authService.isLoggedIn()) {
				this.username = '';
				this.authService.logOut();
				return;
			}
			this.authService.getCurrentUsername()?.subscribe((res: any) => {
				this.username = res.username;
			});
		});
	}
}
