import { Component, Inject, Input } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TGraphGrid } from 'src/app/types/grid-types';
import { Router } from '@angular/router';

const CONSTANTS = {
  Header: 'Grid Map',
};

@Component({
  selector: 'grid-map',
  templateUrl: './grid-map.component.html',
  styleUrls: ['./grid-map.component.scss'],
  providers: [{ provide: 'CONSTANTS', useValue: CONSTANTS }],
})
export class GridMapComponent {
  @Input() data: TGraphGrid | null = null;
  @Input() selectedId: number | undefined;

  public get graph(): TGraphGrid {
    return this.data ?? { nodes: [], links: [] };
  }

  constructor(
    private router: Router,
    @Inject('CONSTANTS') public Constants: any
  ) {}

  handleClick(event: any) {
    if (event.id) {
      this.router.navigate(['/home', `${event.id}`]);
    }
  }
}
