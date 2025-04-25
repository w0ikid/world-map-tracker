import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service'; // Adjust path as needed
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';


@Component({
  imports: [CommonModule, FormsModule],
  selector: 'app-profile',
  template: `
    <div *ngIf="profileData; else loading">
      <h2>Profile</h2>
      <p><strong>Username:</strong> {{ profileData.username }}</p>
      <button (click)="logout()">Logout</button>
    </div>
    <ng-template #loading>
      <p>Loading profile...</p>
    </ng-template>
  `,
  styles: [`
    div {
      padding: 20px;
      max-width: 400px;
      margin: 0 auto;
    }
    button {
      padding: 8px 16px;
      background-color: #dc3545;
      color: white;
      border: none;
      border-radius: 4px;
      cursor: pointer;
    }
    button:hover {
      background-color: #c82333;
    }
  `]
})
export class ProfileComponent implements OnInit {
  profileData: any = null;

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {
    this.authService.profile().subscribe({
      next: (data) => {
        this.profileData = data;
      },
      error: (err) => {
        console.error('Profile fetch error:', err);
        // Redirect to login if unauthorized (e.g., session expired)
        if (err.status === 401) {
          this.router.navigate(['/login']);
        }
      }
    });
  }

  logout(): void {
    this.authService.logout().subscribe({
      next: () => {
        this.router.navigate(['/login']);
      },
      error: (err) => {
        console.error('Logout error:', err);
        // Even if logout fails, redirect to login
        this.router.navigate(['/login']);
      }
    });
  }
}