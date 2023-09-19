import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GridGraphComponent } from './grid-graph.component';

describe('GridGraphComponent', () => {
  let component: GridGraphComponent;
  let fixture: ComponentFixture<GridGraphComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GridGraphComponent]
    });
    fixture = TestBed.createComponent(GridGraphComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
