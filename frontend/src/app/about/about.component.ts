import { Component, OnInit, AfterViewInit, ElementRef, ViewChildren, QueryList, Renderer2, HostListener } from '@angular/core';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements OnInit, AfterViewInit {

  @ViewChildren('contentSection') contentSections!: QueryList<ElementRef>;

  constructor(private renderer: Renderer2, private el: ElementRef) { }

  ngOnInit(): void {
    // Initial setup if needed
  }

  ngAfterViewInit(): void {
    // Generate floating elements right after view initialization
    this.generateFloatingElements(15);
    
    // Check sections initially after a small delay to ensure DOM is ready
    setTimeout(() => {
      this.checkSectionsInView();
    }, 300);
  }

  @HostListener('window:scroll')
  onWindowScroll() {
    this.checkSectionsInView();
  }

  checkSectionsInView() {
    if (this.contentSections && this.contentSections.length) {
      this.contentSections.forEach(section => {
        const rect = section.nativeElement.getBoundingClientRect();
        const windowHeight = window.innerHeight;
        
        // Add active class when section is in viewport
        if (rect.top < windowHeight * 0.75) {
          this.renderer.addClass(section.nativeElement, 'active');
        }
      });
    }
  }

  generateFloatingElements(count: number) {
    // Find the container directly using the native element
    const container = this.el.nativeElement.querySelector('.floating-elements');
    
    if (container) {
      for (let i = 0; i < count; i++) {
        const element = this.renderer.createElement('div');
        this.renderer.addClass(element, 'floating-element');
        
        // Set random properties for each floating element
        this.renderer.setStyle(element, 'top', `${Math.random() * 100}%`);
        this.renderer.setStyle(element, 'left', `${Math.random() * 100}%`);
        
        const size = 20 + Math.random() * 30;
        this.renderer.setStyle(element, 'width', `${size}px`);
        this.renderer.setStyle(element, 'height', `${size}px`);
        
        // Random animation settings
        this.renderer.setStyle(element, 'opacity', '0.1');
        this.renderer.setStyle(element, 'animationDelay', `${Math.random() * 15}s`);
        this.renderer.setStyle(element, 'animationDuration', `${15 + Math.random() * 30}s`);
        
        this.renderer.appendChild(container, element);
      }
    } else {
      console.warn('Floating elements container not found.');
    }
  }
}