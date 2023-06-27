import { Component } from '@angular/core';
import { UserService } from '../shared/user.service';
import { UserData } from '../shared/models';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
	styleUrls: ['./home.component.css'],
})
export class HomeComponent {
	constructor(private userService: UserService) {}

	ngOnInit(): void {
		this.userService
			.getUserData()
			.subscribe((data: UserData) => console.log(data));
	}
}
