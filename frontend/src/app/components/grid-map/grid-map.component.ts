import { Component, Inject, Input } from '@angular/core';
import { TGraphGrid } from 'src/app/types/grid-types';

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
  @Input() data: TGraphGrid = { nodes: [], links: [] };

  constructor(@Inject('CONSTANTS') public Constants: any) {}
}
