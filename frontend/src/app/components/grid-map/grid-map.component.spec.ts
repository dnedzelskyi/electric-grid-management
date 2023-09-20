import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GridMapComponent } from './grid-map.component';

describe('GridMapComponent', () => {
  let component: GridMapComponent;
  let fixture: ComponentFixture<GridMapComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GridMapComponent]
    });
    fixture = TestBed.createComponent(GridMapComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
