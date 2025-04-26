import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service'; // Adjust path as needed
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { CountryStatusService } from '../country-status.service';
import { UserProfile } from '../auth.service'; // Adjust path as needed
import { RouterLink } from '@angular/router';
import { Subscription } from 'rxjs';
import { OnDestroy } from '@angular/core';
@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {
  profileData: UserProfile | null = null;
  visitedPercentage: number = 0;
  visitedCount: number = 0;
  private subscriptions: Subscription[] = [];

  constructor(
    private authService: AuthService,
    private countryStatusService: CountryStatusService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.subscriptions.push(
      this.authService.profile().subscribe({
        next: (data) => {
          this.profileData = data;
        },
        error: (err) => {
          console.error('Profile fetch error:', err);
          if (err.status === 401) {
            this.router.navigate(['/login']);
          }
        }
      })
    );

    this.subscriptions.push(
      this.countryStatusService.getVisitedPercentage().subscribe({
        next: (visitedData) => {
          this.visitedPercentage = visitedData.visited_percentage;
        },
        error: (err) => {
          console.error('Ошибка получения процента посещения стран:', err);
        }
      })
    );

    this.subscriptions.push(
      this.countryStatusService.getVisitedCount().subscribe({
        next: (countData) => {
          this.visitedCount = countData.visited_count;
        },
        error: (err) => {
          console.error('Ошибка получения количества посещенных стран:', err);
        }
      })
    );
  }

  ngOnDestroy(): void {
    // Отписываемся от всех подписок
    this.subscriptions.forEach(sub => sub.unsubscribe());
  }
}