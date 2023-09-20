import { Component, Inject, OnInit } from '@angular/core';
import { HelloServiceService } from '../../services/hello-service.service';
import { GridNode, NodeLink } from '../../pb/grid_pb';
import { ActivatedRoute } from '@angular/router';
import { GridService } from 'src/app/services/grid/grid.service';
import { TApiGrid } from 'src/app/types/grid-types';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss'],
})
export class HomePageComponent implements OnInit {
  selectedNodeId: number | undefined;
  grid: TApiGrid = { nodes: [], links: [] };

  grid1 = {
    nodes: [
      {
        id: '1',
        label: 'Node A',
      },
      {
        id: '2',
        label: 'Node B',
      },
      {
        id: '3',
        label: 'Node C',
      },
      {
        id: '4',
        label: 'Node D',
      },
      {
        id: '5',
        label: 'Node E',
      },
      {
        id: '6',
        label: 'Node F',
      },
    ],
    links: [
      {
        source: '1',
        target: '2',
      },
      {
        source: '1',
        target: '3',
      },
      {
        source: '3',
        target: '4',
      },
      {
        source: '3',
        target: '5',
      },
      {
        source: '4',
        target: '5',
      },
      {
        source: '2',
        target: '6',
      },
    ],
  };

  constructor(
    private gridService: GridService,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    if ('id' in this.route.snapshot.params) {
      this.selectedNodeId = parseInt(this.route.snapshot.params['id']);
    }

    this.gridService.getGrid().then((value) => {
      this.grid = value;
    });
  }
}
