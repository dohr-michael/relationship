import { TestBed, inject } from '@angular/core/testing';

import { UniversesService } from './universes.service';

describe('UniversesService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [UniversesService]
    });
  });

  it('should be created', inject([UniversesService], (service: UniversesService) => {
    expect(service).toBeTruthy();
  }));
});
