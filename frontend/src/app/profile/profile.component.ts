import { Component, OnInit, OnDestroy } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { CountryStatusService } from '../country-status.service';
import { UserProfile } from '../auth.service';
import { Subscription } from 'rxjs';

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
  wishlistCount: number = 0;
  loading: boolean = true;
  private subscriptions: Subscription[] = [];

  constructor(
    private authService: AuthService,
    private countryStatusService: CountryStatusService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.loadProfileData();
  }

  private loadProfileData(): void {
    this.loading = true;
    
    // Get user profile
    this.subscriptions.push(
      this.authService.profile().subscribe({
        next: (data) => {
          this.profileData = data;
          this.loading = false;
        },
        error: (err) => {
          console.error('Profile fetch error:', err);
          if (err.status === 401) {
            this.router.navigate(['/login']);
          }
          this.loading = false;
        }
      })
    );

    // Get visited percentage
    this.subscriptions.push(
      this.countryStatusService.getVisitedPercentage().subscribe({
        next: (visitedData) => {
          this.visitedPercentage = visitedData.visited_percentage;
          // Add animation by starting from 0
          this.animatePercentage(0, this.visitedPercentage);
        },
        error: (err) => {
          console.error('Error getting visited percentage:', err);
        }
      })
    );

    // Get visited count
    this.subscriptions.push(
      this.countryStatusService.getVisitedCount().subscribe({
        next: (countData) => {
          this.visitedCount = countData.visited_count;
        },
        error: (err) => {
          console.error('Error getting visited count:', err);
        }
      })
    );
    
    // Get wishlist count - you would need to add this method to your service
    this.subscriptions.push(
      this.countryStatusService.getWishlistCount().subscribe({
        next: (wishlistData) => {
          this.wishlistCount = wishlistData.wishlist_count;
        },
        error: (err) => {
          console.error('Error getting wishlist count:', err);
        }
      })
    );
  }
  
  // Animation for percentage counter
  private animatePercentage(start: number, end: number): void {
    const duration = 1500;
    const frameDuration = 1000 / 60;
    const totalFrames = Math.round(duration / frameDuration);
    let frame = 0;
    
    const animate = () => {
      frame++;
      const progress = frame / totalFrames;
      const currentValue = Math.round(start + (end - start) * progress);
      
      if (frame === totalFrames) {
        this.visitedPercentage = end;
      } else {
        this.visitedPercentage = currentValue;
        requestAnimationFrame(animate);
      }
    };
    
    requestAnimationFrame(animate);
  }
  
  getTravelStatus(): string {
    if (this.visitedPercentage < 10) return 'Beginner';
    if (this.visitedPercentage < 25) return 'Explorer';
    if (this.visitedPercentage < 50) return 'Adventurer';
    if (this.visitedPercentage < 75) return 'Globetrotter';
    return 'World Master';
  }

  ngOnDestroy(): void {
    this.subscriptions.forEach(sub => sub.unsubscribe());
  }
}