import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-statistics',
  imports: [CommonModule],
  templateUrl: './statistics.component.html',
  styleUrl: './statistics.component.css'
})
export class StatisticsComponent implements OnInit {
  topVisitedCountries: any[] = [];
  topWishlistCountries: any[] = [];
  isLoading = true;
  error: string | null = null;

  constructor(private statisticsService: StatisticsService) { }

  async ngOnInit(): Promise<void> {
    try {
      this.isLoading = true;
      
      // Fetch both statistics in parallel
      const [visitedData, wishlistData] = await Promise.all([
        this.statisticsService.getTopVisitedCountries(),
        this.statisticsService.getTopWishlistCountries()
      ]);
      
      this.topVisitedCountries = visitedData;
      this.topWishlistCountries = wishlistData;
    } catch (err) {
      console.error('Error fetching statistics:', err);
      this.error = 'Failed to load statistics data';
    } finally {
      this.isLoading = false;
    }
  }
}
