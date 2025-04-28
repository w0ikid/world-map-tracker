import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../auth.service';
import { CountryStatusService } from '../country-status.service';
import { UserProfile } from '../auth.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-user-profile',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit, OnDestroy {
  username: string = '';
  profileData: UserProfile | null = null;
  visitedPercentage: number = 0;
  visitedCount: number = 0;
  wishlistCount: number = 0;
  loading: boolean = true;
  error: string | null = null;
  private subscriptions: Subscription[] = [];

  constructor(
    private authService: AuthService,
    private countryStatusService: CountryStatusService,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.subscriptions.push(
      this.route.params.subscribe(params => {
        this.username = params['username'];
        if (this.username) {
          this.loadUserProfileData();
        } else {
          this.error = 'Username not provided';
          this.loading = false;
        }
      })
    );
  }

  private loadUserProfileData(): void {
    this.loading = true;
    this.error = null;
    
    // Get user profile by username
    this.subscriptions.push(
      this.authService.getUserByUsername(this.username).subscribe({
        next: (data) => {
          this.profileData = data;
          this.loading = false;
        },
        error: (err) => {
          console.error('User profile fetch error:', err);
          this.error = `Failed to load profile for ${this.username}`;
          this.loading = false;
        }
      })
    );

    // Get visited percentage by username
    this.subscriptions.push(
      this.countryStatusService.getVisitedPercentageByUsername(this.username).subscribe({
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

    // Get visited count by username
    this.subscriptions.push(
      this.countryStatusService.getVisitedCountByUsername(this.username).subscribe({
        next: (countData) => {
          this.visitedCount = countData.visited_count;
        },
        error: (err) => {
          console.error('Error getting visited count:', err);
        }
      })
    );
    
    // Get wishlist count by username
    this.subscriptions.push(
      this.countryStatusService.getWishlistCountByUsername(this.username).subscribe({
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