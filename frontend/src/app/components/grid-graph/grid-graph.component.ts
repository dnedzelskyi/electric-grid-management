import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import * as d3 from 'd3';
import { SimulationNodeDatum } from 'd3';

@Component({
  selector: 'grid-graph',
  templateUrl: './grid-graph.component.html',
  styleUrls: ['./grid-graph.component.scss'],
})
export class GridGraphComponent implements OnInit {
  @ViewChild('graphSVG', { static: true })
  private graphContainer!: ElementRef;

  constructor() {}

  ngOnInit(): void {
    // Sample graph data (you can replace this with your own data)
    const nodes = [
      { id: 'A', x: 100, y: 100 },
      { id: 'B', x: 400, y: 400 },
      { id: 'C', x: 800, y: 800 },
    ];
    const links = [
      { source: 'A', target: 'B' },
      { source: 'B', target: 'C' },
      { source: 'C', target: 'A' },
    ];

    // Create a D3.js force-directed graph simulation
    const svg = d3.select(this.graphContainer.nativeElement),
      width = +svg.attr('width'),
      height = +svg.attr('height');

    const simulation = d3
      .forceSimulation(nodes)
      .force(
        'link',
        d3.forceLink(links).id((d) => (d as any).id)
      )
      .force('charge', d3.forceManyBody())
      .force('center', d3.forceCenter(width / 2, height / 2));

    // Create links
    const link = svg
      .append('g')
      .selectAll('line')
      .data(links)
      .enter()
      .append('line')
      .attr('stroke-width', 2)
      .attr('stroke', 'black');

    // Create nodes
    const node = svg
      .append('g')
      .selectAll('circle')
      .data(nodes)
      .enter()
      .append('circle')
      .attr('r', 10)
      .attr('fill', 'blue');

    // Add labels to nodes
    const label = svg
      .selectAll('.label')
      .data(nodes)
      .enter()
      .append('text')
      .attr('class', 'label')
      .attr('dy', '.35em')
      .text((d) => d.id);

    // Define tick function for the simulation
    const ticked = () => {
      // link
      //   .attr('x1', (d) => d.source.x)
      //   .attr('y1', (d) => d.source.y)
      //   .attr('x2', (d) => d.target.x)
      //   .attr('y2', (d) => d.target.y);

      node.attr('cx', (d) => d.x + 100).attr('cy', (d) => d.y + 100);

      label.attr('x', (d) => d.x).attr('y', (d) => d.y);
    };

    // Add the tick function to the simulation
    simulation.nodes(nodes).on('tick', ticked);
    simulation.force<d3.ForceLink<any, any>>('link')?.links(links);
  }
}
