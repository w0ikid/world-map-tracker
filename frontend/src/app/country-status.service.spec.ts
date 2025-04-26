import { TestBed } from '@angular/core/testing';

import { CountryStatusService } from './country-status.service';

describe('CountryStatusService', () => {
  let service: CountryStatusService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CountryStatusService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
