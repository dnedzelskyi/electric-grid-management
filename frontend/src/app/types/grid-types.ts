import * as pb from '../pb/grid_pb';
import { Node, Edge } from '@swimlane/ngx-graph';

export type TApiGrid = { nodes: pb.GridNode[]; links: pb.NodeLink[] };
export type TGraphGrid = { nodes: Node[]; links: Edge[] };
