import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GridNodeDetailsComponent } from './grid-node-details.component';

describe('GridNodeDetailsComponent', () => {
  let component: GridNodeDetailsComponent;
  let fixture: ComponentFixture<GridNodeDetailsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GridNodeDetailsComponent]
    });
    fixture = TestBed.createComponent(GridNodeDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
