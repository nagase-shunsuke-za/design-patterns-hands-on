interface ITransport {
    order(): void
    delivery(target: string): void
}

class SeaTransport implements ITransport {
    private id: number;

    constructor(id: number) {
        this.id = id;
    }

    order(): void {
        console.log(`SeaTransport ID: ${this.id}: 配送受注`);
    }
    delivery(target: string): void {
        console.log(`SeaTransport ID: ${this.id}: ${target}配送実施`);
    }
}

class RoadTransport implements ITransport {
    private id: number;

    constructor(id: number) {
        this.id = id;
    }

    order(): void {
        console.log(`RoadTransport ID: ${this.id}: 配送受注`);
    }
    delivery(target: string): void {
        console.log(`RoadTransport ID: ${this.id}: ${target}配送実施`);
    }
}

export {ITransport, SeaTransport, RoadTransport}
